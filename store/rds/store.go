package rds

import (
	"context"
	"database/sql"
	"github.com/usememos/memos/api"
	"github.com/usememos/memos/store"
	"sync"

	"github.com/usememos/memos/server/profile"
)

// Store provides database access to all raw objects.
type Store struct {
	db      *sql.DB
	profile *profile.Profile

	userCache        sync.Map // map[int]*userRaw
	userSettingCache sync.Map // map[string]*userSettingRaw
	memoCache        sync.Map // map[int]*memoRaw
	shortcutCache    sync.Map // map[int]*shortcutRaw
	idpCache         sync.Map // map[int]*identityProviderMessage
}

// New creates a new instance of Store.
func New(db *sql.DB, profile *profile.Profile) *Store {
	return &Store{
		db:      db,
		profile: profile,
	}
}

func (s *Store) CreateIdentityProvider(ctx context.Context, create *store.IdentityProviderMessage) (*store.IdentityProviderMessage, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Store) ListIdentityProviders(ctx context.Context, find *store.FindIdentityProviderMessage) ([]*store.IdentityProviderMessage, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Store) GetIdentityProvider(ctx context.Context, find *store.FindIdentityProviderMessage) (*store.IdentityProviderMessage, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Store) UpdateIdentityProvider(ctx context.Context, update *store.UpdateIdentityProviderMessage) (*store.IdentityProviderMessage, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Store) DeleteIdentityProvider(ctx context.Context, delete *store.DeleteIdentityProviderMessage) error {
	// TODO implement me
	panic("implement me")
}

func (s *Store) UpsertTag(ctx context.Context, upsert *api.TagUpsert) (*api.Tag, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Store) FindTagList(ctx context.Context, find *api.TagFind) ([]*api.Tag, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Store) DeleteTag(ctx context.Context, delete *api.TagDelete) error {
	// TODO implement me
	panic("implement me")
}

func (s *Store) Vacuum(ctx context.Context) error {
	// TODO implement me
	panic("implement me")
}

func (s *Store) UpsertUserSetting(ctx context.Context, upsert *api.UserSettingUpsert) (*api.UserSetting, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Store) FindUserSettingList(ctx context.Context, find *api.UserSettingFind) ([]*api.UserSetting, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Store) FindUserSetting(ctx context.Context, find *api.UserSettingFind) (*api.UserSetting, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Store) ComposeMemo(ctx context.Context, memo *api.Memo) (*api.Memo, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Store) CreateMemo(ctx context.Context, create *api.MemoCreate) (*api.Memo, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Store) PatchMemo(ctx context.Context, patch *api.MemoPatch) (*api.Memo, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Store) FindMemoList(ctx context.Context, find *api.MemoFind) ([]*api.Memo, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Store) FindMemo(ctx context.Context, find *api.MemoFind) (*api.Memo, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Store) DeleteMemo(ctx context.Context, delete *api.MemoDelete) error {
	// TODO implement me
	panic("implement me")
}

func (s *Store) FindMemoResourceList(ctx context.Context, find *api.MemoResourceFind) ([]*api.MemoResource, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Store) FindMemoResource(ctx context.Context, find *api.MemoResourceFind) (*api.MemoResource, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Store) UpsertMemoResource(ctx context.Context, upsert *api.MemoResourceUpsert) (*api.MemoResource, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Store) DeleteMemoResource(ctx context.Context, delete *api.MemoResourceDelete) error {
	// TODO implement me
	panic("implement me")
}

func (s *Store) UpsertSystemSetting(ctx context.Context, upsert *api.SystemSettingUpsert) (*api.SystemSetting, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Store) FindSystemSettingList(ctx context.Context, find *api.SystemSettingFind) ([]*api.SystemSetting, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Store) FindSystemSetting(ctx context.Context, find *api.SystemSettingFind) (*api.SystemSetting, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Store) ComposeMemoCreator(ctx context.Context, memo *api.Memo) error {
	// TODO implement me
	panic("implement me")
}

func (s *Store) CreateUser(ctx context.Context, create *api.UserCreate) (*api.User, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Store) PatchUser(ctx context.Context, patch *api.UserPatch) (*api.User, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Store) FindUserList(ctx context.Context, find *api.UserFind) ([]*api.User, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Store) FindUser(ctx context.Context, find *api.UserFind) (*api.User, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Store) DeleteUser(ctx context.Context, delete *api.UserDelete) error {
	// TODO implement me
	panic("implement me")
}

func (s *Store) CreateStorage(ctx context.Context, create *api.StorageCreate) (*api.Storage, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Store) PatchStorage(ctx context.Context, patch *api.StoragePatch) (*api.Storage, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Store) FindStorageList(ctx context.Context, find *api.StorageFind) ([]*api.Storage, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Store) FindStorage(ctx context.Context, find *api.StorageFind) (*api.Storage, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Store) DeleteStorage(ctx context.Context, delete *api.StorageDelete) error {
	// TODO implement me
	panic("implement me")
}

func (s *Store) CreateActivity(ctx context.Context, create *api.ActivityCreate) (*api.Activity, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Store) ComposeMemoResourceList(ctx context.Context, memo *api.Memo) error {
	// TODO implement me
	panic("implement me")
}

func (s *Store) CreateResource(ctx context.Context, create *api.ResourceCreate) (*api.Resource, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Store) FindResourceList(ctx context.Context, find *api.ResourceFind) ([]*api.Resource, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Store) FindResource(ctx context.Context, find *api.ResourceFind) (*api.Resource, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Store) DeleteResource(ctx context.Context, delete *api.ResourceDelete) error {
	// TODO implement me
	panic("implement me")
}

func (s *Store) PatchResource(ctx context.Context, patch *api.ResourcePatch) (*api.Resource, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Store) FindMemoOrganizer(ctx context.Context, find *api.MemoOrganizerFind) (*api.MemoOrganizer, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Store) UpsertMemoOrganizer(ctx context.Context, upsert *api.MemoOrganizerUpsert) error {
	// TODO implement me
	panic("implement me")
}

func (s *Store) DeleteMemoOrganizer(ctx context.Context, delete *api.MemoOrganizerDelete) error {
	// TODO implement me
	panic("implement me")
}

func (s *Store) CreateShortcut(ctx context.Context, create *api.ShortcutCreate) (*api.Shortcut, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Store) PatchShortcut(ctx context.Context, patch *api.ShortcutPatch) (*api.Shortcut, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Store) FindShortcutList(ctx context.Context, find *api.ShortcutFind) ([]*api.Shortcut, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Store) FindShortcut(ctx context.Context, find *api.ShortcutFind) (*api.Shortcut, error) {
	// TODO implement me
	panic("implement me")
}

func (s *Store) DeleteShortcut(ctx context.Context, delete *api.ShortcutDelete) error {
	// TODO implement me
	panic("implement me")
}
