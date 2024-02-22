// Code generated by interfacer; DO NOT EDIT

package agent

import (
	"context"
	"github.com/buildkite/agent/v3/api"
)

// APIClient is an interface generated for "github.com/buildkite/agent/v3/api.Client".
type APIClient interface {
	AcceptJob(context.Context, *api.Job) (*api.Job, *api.Response, error)
	AcquireJob(context.Context, string, ...api.Header) (*api.Job, *api.Response, error)
	Annotate(context.Context, string, *api.Annotation) (*api.Response, error)
	AnnotationRemove(context.Context, string, string) (*api.Response, error)
	Config() api.Config
	Connect(context.Context) (*api.Response, error)
	CreateArtifacts(context.Context, string, *api.ArtifactBatch) (*api.ArtifactBatchCreateResponse, *api.Response, error)
	Disconnect(context.Context) (*api.Response, error)
	ExistsMetaData(context.Context, string, string, string) (*api.MetaDataExists, *api.Response, error)
	FinishJob(context.Context, *api.Job) (*api.Response, error)
	FromAgentRegisterResponse(*api.AgentRegisterResponse) *api.Client
	FromPing(*api.Ping) *api.Client
	GenerateGithubCodeAccessToken(context.Context, string, string) (string, *api.Response, error)
	GetJobState(context.Context, string) (*api.JobState, *api.Response, error)
	GetMetaData(context.Context, string, string, string) (*api.MetaData, *api.Response, error)
	GetSecret(context.Context, *api.GetSecretRequest) (*api.Secret, *api.Response, error)
	Heartbeat(context.Context) (*api.Heartbeat, *api.Response, error)
	MetaDataKeys(context.Context, string, string) ([]string, *api.Response, error)
	OIDCToken(context.Context, *api.OIDCTokenRequest) (*api.OIDCToken, *api.Response, error)
	Ping(context.Context) (*api.Ping, *api.Response, error)
	PipelineUploadStatus(context.Context, string, string, ...api.Header) (*api.PipelineUploadStatus, *api.Response, error)
	Register(context.Context, *api.AgentRegisterRequest) (*api.AgentRegisterResponse, *api.Response, error)
	SaveHeaderTimes(context.Context, string, *api.HeaderTimes) (*api.Response, error)
	SearchArtifacts(context.Context, string, *api.ArtifactSearchOptions) ([]*api.Artifact, *api.Response, error)
	SetMetaData(context.Context, string, *api.MetaData) (*api.Response, error)
	StartJob(context.Context, *api.Job) (*api.Response, error)
	StepExport(context.Context, string, *api.StepExportRequest) (*api.StepExportResponse, *api.Response, error)
	StepUpdate(context.Context, string, *api.StepUpdate) (*api.Response, error)
	UpdateArtifacts(context.Context, string, map[string]string) (*api.Response, error)
	UploadChunk(context.Context, string, *api.Chunk) (*api.Response, error)
	UploadPipeline(context.Context, string, *api.PipelineChange, ...api.Header) (*api.Response, error)
}
