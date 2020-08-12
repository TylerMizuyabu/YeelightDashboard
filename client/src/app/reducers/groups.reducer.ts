import { createReducer, on } from '@ngrx/store';
import { newGroup } from '../actions/new-group.action';
import {Yeelight} from '../models/light';

export const groupsReducer = createReducer(
  new Map<string, Set<Yeelight>>(),
  on(newGroup, (state, {groupName}) => (new Map<string, Set<Yeelight>>([...state, [groupName, new Set<Yeelight>()]])))
);