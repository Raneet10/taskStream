package taskstreamer

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/bank"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/raneet10/taskStream/x/taskstreamer/types"
)

type Keeper struct {
	bank bank.Keeper

	taskStore sdk.StoreKey

	cdc *codec.Codec
}

func NewKeeper(bank bank.Keeper, taskStore sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		bank:      bank,
		taskStore: taskStore,
		cdc:       cdc,
	}
}

// Get a individual task
func (k Keeper) GetTask(ctx sdk.Context, key string) (types.Task, error) {
	store := sdk.KVStrore(k.taskStore)
	var task types.Task
	byteKey := []byte(types.TaskPrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &task)

	if err != nil {
		return task, err
	}
	return task, nil
}

// Get all tasks
func (k Keeper) GetAllTasks(ctx sdk.Context) ([]types.Task, error) {
	store := sdk.KVStore(k.taskStore)
	var tasks []types.Task
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.TaskPrefix))

	for ; iterator.Valid(); iterator.Next() {
		var task types.Task
		err := k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &task)
		if err != nil {
			return tasks, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

// CreateTask
// Add value to task - add new value giver as backer
// Give proof of compeltion
// Vote if the prof is accepted
// Payout the person who completed it
