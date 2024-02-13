package sync

import (
	"fmt"
	"sort"
	"time"

	"github.com/Sceptyre/go-five9-scim/internal/logger"
	"github.com/Sceptyre/go-five9-scim/internal/mappers"
	"github.com/Sceptyre/go-five9-scim/pkg/five9/api"

	"github.com/elimity-com/scim"
)

var SyncData map[string]scim.Resource
var SyncIdsToUsernames map[string]string
var Logger = logger.Logger{
	Namespace: "FIVE9_SYNC",
}

func SyncFive9() {
	Logger.Log("Syncing Five9 DB Data")

	u, err := api.GetUsersInfo(".*")

	if err != nil {
		Logger.LogError(string(err.Body.Error.Message))
		return
	}

	Logger.Log(fmt.Sprintf("Found %v users", len(u.Return)))

	syncIdMap := map[string]string{}
	for _, v := range u.Return {
		syncIdMap[fmt.Sprint(v.GeneralInfo.Id)] = v.GeneralInfo.UserName
		mappers.MapFive9UserInfoToScimUser(v)
	}
	SyncIdsToUsernames = syncIdMap

	scimUsers := mappers.MapFiveUserInfoListToScimUserList(&u.Return)

	scimUsersMap := map[string]scim.Resource{}
	for _, v := range scimUsers {
		scimUsersMap[fmt.Sprint(v.ID)] = v
	}
	SyncData = scimUsersMap

	Logger.Log("Sync Complete")
}

func backgroundSyncFive9() {
	for {
		time.Sleep(15 * time.Minute)
		SyncFive9()
	}
}

func GetMapValues[T any](sourceMap map[string]T) []T {
	output := []T{}
	keys := []string{}

	for k := range sourceMap {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		output = append(output, sourceMap[k])
	}

	return output
}

func Sync() {
	go backgroundSyncFive9()
}
