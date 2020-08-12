import {
  ActionReducer,
  ActionReducerMap,
  createFeatureSelector,
  createSelector,
  MetaReducer,
  createReducer
} from '@ngrx/store';
import { environment } from '../../environments/environment';
import {Yeelight} from '../models/light';
import {lightsReducer} from './lights.reducer';
import {groupsReducer} from './groups.reducer';

export interface State {
  lights: Set<Yeelight>;
  groups: Map<string, Set<Yeelight>>;
}

export const reducers: ActionReducerMap<State> = {
  lights: lightsReducer,
  groups: groupsReducer,
};


export const metaReducers: MetaReducer<State>[] = !environment.production ? [] : [];
