import { createReducer, on } from '@ngrx/store';
import { newGroup } from '../actions/new-group.action';
import {Yeelight} from '../models/light';
import {Group} from '../models/group';
import { EntityState, createEntityAdapter } from '@ngrx/entity';

export interface GroupsState extends EntityState<Group> {}

const groupsAdapter = createEntityAdapter<Group>();

export const initialGroupsState = groupsAdapter.getInitialState({
  selectId: (g: Group) => (g.name),
  sortComparer: (a: Group, b: Group) => (a.name.localeCompare(b.name))
});

export const groupsReducer = createReducer(
  initialGroupsState,
  on(newGroup, (state, {groupName}) => groupsAdapter.addOne({name: groupName, lights: new Set<Yeelight>()}, state))
);