package repair

import (
	"context"
	"errors"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx"
	"github.com/scylladb/gocqlx/qb"
	"github.com/scylladb/mermaid"
	"github.com/scylladb/mermaid/log"
	"github.com/scylladb/mermaid/schema"
)

// globalClusterID is a special value used as a cluster ID for a global
// configuration.
var globalClusterID = mermaid.UUID{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x40, 0x0, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}

// Service orchestrates cluster repairs.
type Service struct {
	session *gocql.Session
	logger  log.Logger
}

// NewService creates a new service instance.
func NewService(session *gocql.Session, logger log.Logger) (*Service, error) {
	if session == nil || session.Closed() {
		return nil, errors.New("nil or closed session")
	}

	return &Service{
		session: session,
		logger:  logger,
	}, nil
}

// GetConfig returns repair configuration for a given object. If nothing
// was found mermaid.ErrNotFound is returned.
func (s *Service) GetConfig(ctx context.Context, clusterID mermaid.UUID, t ConfigType, externalID string) (*Config, error) {
	s.logger.Debug(ctx, "GetConfig",
		"cluster_id", clusterID,
		"type", t,
		"external_id", externalID,
	)

	stmt, names := schema.RepairConfig.Get()
	q := gocqlx.Query(s.session.Query(stmt).WithContext(ctx), names).BindMap(qb.M{
		"cluster_id":  clusterID,
		"type":        t,
		"external_id": externalID,
	})

	var c Config
	if err := gocqlx.Iter(q.Query).Unsafe().Get(&c); err != nil {
		return nil, err
	}

	return &c, nil
}

// PutConfig upserts repair configuration for a given object. If c is nil
// the configuration is removed.
func (s *Service) PutConfig(ctx context.Context, clusterID mermaid.UUID, t ConfigType, externalID string, c *Config) error {
	s.logger.Debug(ctx, "PutConfig",
		"cluster_id", clusterID,
		"type", t,
		"external_id", externalID,
		"config", c,
	)

	if err := c.Validate(); err != nil {
		return err
	}

	stmt, names := schema.RepairConfig.Insert()
	return gocqlx.Query(s.session.Query(stmt).WithContext(ctx), names).BindStructMap(c, qb.M{
		"cluster_id":  clusterID,
		"type":        t,
		"external_id": externalID,
	}).ExecRelease()
}

// DeleteConfig removes repair configuration for a given object.
func (s *Service) DeleteConfig(ctx context.Context, clusterID mermaid.UUID, t ConfigType, externalID string) error {
	s.logger.Debug(ctx, "DeleteConfig",
		"cluster_id", clusterID,
		"type", t,
		"external_id", externalID,
	)

	stmt, names := schema.RepairConfig.Delete()
	return gocqlx.Query(s.session.Query(stmt).WithContext(ctx), names).BindMap(qb.M{
		"cluster_id":  clusterID,
		"type":        t,
		"external_id": externalID,
	}).ExecRelease()
}
