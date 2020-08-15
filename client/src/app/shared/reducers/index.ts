import {
  ActionReducerMap,
  MetaReducer
} from '@ngrx/store';
import {EntityState} from '@ngrx/entity';
import {lightsReducer, initialLightsState, LightsState} from './lights.reducer';
import {groupsReducer, initialGroupsState, GroupsState} from './groups.reducer';
import {environment} from '../../../environments/environment';

export interface State {
  lights: LightsState;
  groups: GroupsState;
}

export const initialState = {
  lights: initialLightsState,
  groups: initialGroupsState
}

export const reducers: ActionReducerMap<State> = {
  lights: lightsReducer,
  groups: groupsReducer,
};


export const metaReducers: MetaReducer<State>[] = !environment.production ? [] : [];
