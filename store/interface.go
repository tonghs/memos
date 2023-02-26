package store

import (
	"context"
	"database/sql"

	"github.com/usememos/memos/api"
)

const (
	MySQLDriver    = "mysql"
	PostgresDriver = "postgres"
	SQLiteDriver   = "sqlite"
)

type DB interface {
	Open(ctx context.Context) (err error)
	Instance() *sql.DB
}

type Store interface {
	CreateIdentityProvider(ctx context.Context, create *IdentityProviderMessage) (*IdentityProviderMessage, error)
	ListIdentityProviders(ctx context.Context, find *FindIdentityProviderMessage) ([]*IdentityProviderMessage, error)
	GetIdentityProvider(ctx context.Context, find *FindIdentityProviderMessage) (*IdentityProviderMessage, error)
	UpdateIdentityProvider(ctx context.Context, update *UpdateIdentityProviderMessage) (*IdentityProviderMessage, error)
	DeleteIdentityProvider(ctx context.Context, delete *DeleteIdentityProviderMessage) error
	UpsertTag(ctx context.Context, upsert *api.TagUpsert) (*api.Tag, error)
	FindTagList(ctx context.Context, find *api.TagFind) ([]*api.Tag, error)
	DeleteTag(ctx context.Context, delete *api.TagDelete) error
	Vacuum(ctx context.Context) error
	UpsertUserSetting(ctx context.Context, upsert *api.UserSettingUpsert) (*api.UserSetting, error)
	FindUserSettingList(ctx context.Context, find *api.UserSettingFind) ([]*api.UserSetting, error)
	FindUserSetting(ctx context.Context, find *api.UserSettingFind) (*api.UserSetting, error)
	ComposeMemo(ctx context.Context, memo *api.Memo) (*api.Memo, error)
	CreateMemo(ctx context.Context, create *api.MemoCreate) (*api.Memo, error)
	PatchMemo(ctx context.Context, patch *api.MemoPatch) (*api.Memo, error)
	FindMemoList(ctx context.Context, find *api.MemoFind) ([]*api.Memo, error)
	FindMemo(ctx context.Context, find *api.MemoFind) (*api.Memo, error)
	DeleteMemo(ctx context.Context, delete *api.MemoDelete) error
	FindMemoResourceList(ctx context.Context, find *api.MemoResourceFind) ([]*api.MemoResource, error)
	FindMemoResource(ctx context.Context, find *api.MemoResourceFind) (*api.MemoResource, error)
	UpsertMemoResource(ctx context.Context, upsert *api.MemoResourceUpsert) (*api.MemoResource, error)
	DeleteMemoResource(ctx context.Context, delete *api.MemoResourceDelete) error
	UpsertSystemSetting(ctx context.Context, upsert *api.SystemSettingUpsert) (*api.SystemSetting, error)
	FindSystemSettingList(ctx context.Context, find *api.SystemSettingFind) ([]*api.SystemSetting, error)
	FindSystemSetting(ctx context.Context, find *api.SystemSettingFind) (*api.SystemSetting, error)
	ComposeMemoCreator(ctx context.Context, memo *api.Memo) error
	CreateUser(ctx context.Context, create *api.UserCreate) (*api.User, error)
	PatchUser(ctx context.Context, patch *api.UserPatch) (*api.User, error)
	FindUserList(ctx context.Context, find *api.UserFind) ([]*api.User, error)
	FindUser(ctx context.Context, find *api.UserFind) (*api.User, error)
	DeleteUser(ctx context.Context, delete *api.UserDelete) error
	CreateStorage(ctx context.Context, create *api.StorageCreate) (*api.Storage, error)
	PatchStorage(ctx context.Context, patch *api.StoragePatch) (*api.Storage, error)
	FindStorageList(ctx context.Context, find *api.StorageFind) ([]*api.Storage, error)
	FindStorage(ctx context.Context, find *api.StorageFind) (*api.Storage, error)
	DeleteStorage(ctx context.Context, delete *api.StorageDelete) error
	CreateActivity(ctx context.Context, create *api.ActivityCreate) (*api.Activity, error)
	ComposeMemoResourceList(ctx context.Context, memo *api.Memo) error
	CreateResource(ctx context.Context, create *api.ResourceCreate) (*api.Resource, error)
	FindResourceList(ctx context.Context, find *api.ResourceFind) ([]*api.Resource, error)
	FindResource(ctx context.Context, find *api.ResourceFind) (*api.Resource, error)
	DeleteResource(ctx context.Context, delete *api.ResourceDelete) error
	PatchResource(ctx context.Context, patch *api.ResourcePatch) (*api.Resource, error)
	FindMemoOrganizer(ctx context.Context, find *api.MemoOrganizerFind) (*api.MemoOrganizer, error)
	UpsertMemoOrganizer(ctx context.Context, upsert *api.MemoOrganizerUpsert) error
	DeleteMemoOrganizer(ctx context.Context, delete *api.MemoOrganizerDelete) error
	CreateShortcut(ctx context.Context, create *api.ShortcutCreate) (*api.Shortcut, error)
	PatchShortcut(ctx context.Context, patch *api.ShortcutPatch) (*api.Shortcut, error)
	FindShortcutList(ctx context.Context, find *api.ShortcutFind) ([]*api.Shortcut, error)
	FindShortcut(ctx context.Context, find *api.ShortcutFind) (*api.Shortcut, error)
	DeleteShortcut(ctx context.Context, delete *api.ShortcutDelete) error
}
