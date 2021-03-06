// Copyright (c) 2018-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package store

import (
	"context"

	"github.com/crspeller/mattermost-server/model"
)

func (s *LocalCacheSupplier) handleClusterInvalidateGroup(msg *model.ClusterMessage) {
	if msg.Data == CLEAR_CACHE_MESSAGE_DATA {
		s.groupCache.Purge()
	} else {
		s.groupCache.Remove(msg.Data)
	}
}

func (s *LocalCacheSupplier) GroupCreate(ctx context.Context, group *model.Group, hints ...LayeredStoreHint) *LayeredStoreSupplierResult {
	return s.Next().GroupCreate(ctx, group, hints...)
}

func (s *LocalCacheSupplier) GroupGet(ctx context.Context, groupID string, hints ...LayeredStoreHint) *LayeredStoreSupplierResult {
	if result := s.doStandardReadCache(ctx, s.groupCache, groupID, hints...); result != nil {
		return result
	}

	result := s.Next().GroupGet(ctx, groupID, hints...)

	s.doStandardAddToCache(ctx, s.groupCache, groupID, result, hints...)

	return result
}

func (s *LocalCacheSupplier) GroupGetByRemoteID(ctx context.Context, remoteID string, groupSource model.GroupSource, hints ...LayeredStoreHint) *LayeredStoreSupplierResult {
	return s.Next().GroupGetByRemoteID(ctx, remoteID, groupSource, hints...)
}

func (s *LocalCacheSupplier) GroupGetAllBySource(ctx context.Context, groupSource model.GroupSource, hints ...LayeredStoreHint) *LayeredStoreSupplierResult {
	return s.Next().GroupGetAllBySource(ctx, groupSource, hints...)
}

func (s *LocalCacheSupplier) GroupUpdate(ctx context.Context, group *model.Group, hints ...LayeredStoreHint) *LayeredStoreSupplierResult {
	defer s.doInvalidateCacheCluster(s.groupCache, group.Id)
	return s.Next().GroupUpdate(ctx, group, hints...)
}

func (s *LocalCacheSupplier) GroupDelete(ctx context.Context, groupID string, hints ...LayeredStoreHint) *LayeredStoreSupplierResult {
	defer s.doInvalidateCacheCluster(s.groupCache, groupID)
	defer s.doClearCacheCluster(s.groupCache)

	return s.Next().GroupDelete(ctx, groupID, hints...)
}

func (s *LocalCacheSupplier) GroupGetMemberUsers(ctx context.Context, groupID string, hints ...LayeredStoreHint) *LayeredStoreSupplierResult {
	return s.Next().GroupGetMemberUsers(ctx, groupID, hints...)
}

func (s *LocalCacheSupplier) GroupGetMemberUsersPage(ctx context.Context, groupID string, offset int, limit int, hints ...LayeredStoreHint) *LayeredStoreSupplierResult {
	return s.Next().GroupGetMemberUsersPage(ctx, groupID, offset, limit, hints...)
}

func (s *LocalCacheSupplier) GroupGetMemberCount(ctx context.Context, groupID string, hints ...LayeredStoreHint) *LayeredStoreSupplierResult {
	return s.Next().GroupGetMemberCount(ctx, groupID, hints...)
}

func (s *LocalCacheSupplier) GroupCreateOrRestoreMember(ctx context.Context, groupID string, userID string, hints ...LayeredStoreHint) *LayeredStoreSupplierResult {
	return s.Next().GroupCreateOrRestoreMember(ctx, groupID, userID, hints...)
}

func (s *LocalCacheSupplier) GroupDeleteMember(ctx context.Context, groupID string, userID string, hints ...LayeredStoreHint) *LayeredStoreSupplierResult {
	return s.Next().GroupDeleteMember(ctx, groupID, userID, hints...)
}

func (s *LocalCacheSupplier) GroupCreateGroupSyncable(ctx context.Context, groupSyncable *model.GroupSyncable, hints ...LayeredStoreHint) *LayeredStoreSupplierResult {
	return s.Next().GroupCreateGroupSyncable(ctx, groupSyncable, hints...)
}

func (s *LocalCacheSupplier) GroupGetGroupSyncable(ctx context.Context, groupID string, syncableID string, syncableType model.GroupSyncableType, hints ...LayeredStoreHint) *LayeredStoreSupplierResult {
	return s.Next().GroupGetGroupSyncable(ctx, groupID, syncableID, syncableType, hints...)
}

func (s *LocalCacheSupplier) GroupGetAllGroupSyncablesByGroup(ctx context.Context, groupID string, syncableType model.GroupSyncableType, hints ...LayeredStoreHint) *LayeredStoreSupplierResult {
	return s.Next().GroupGetAllGroupSyncablesByGroup(ctx, groupID, syncableType, hints...)
}

func (s *LocalCacheSupplier) GroupUpdateGroupSyncable(ctx context.Context, groupSyncable *model.GroupSyncable, hints ...LayeredStoreHint) *LayeredStoreSupplierResult {
	return s.Next().GroupUpdateGroupSyncable(ctx, groupSyncable, hints...)
}

func (s *LocalCacheSupplier) GroupDeleteGroupSyncable(ctx context.Context, groupID string, syncableID string, syncableType model.GroupSyncableType, hints ...LayeredStoreHint) *LayeredStoreSupplierResult {
	return s.Next().GroupDeleteGroupSyncable(ctx, groupID, syncableID, syncableType, hints...)
}

func (s *LocalCacheSupplier) PendingAutoAddTeamMembers(ctx context.Context, minGroupMembersCreateAt int64, hints ...LayeredStoreHint) *LayeredStoreSupplierResult {
	return s.Next().PendingAutoAddTeamMembers(ctx, minGroupMembersCreateAt, hints...)
}

func (s *LocalCacheSupplier) PendingAutoAddChannelMembers(ctx context.Context, minGroupMembersCreateAt int64, hints ...LayeredStoreHint) *LayeredStoreSupplierResult {
	return s.Next().PendingAutoAddChannelMembers(ctx, minGroupMembersCreateAt, hints...)
}
