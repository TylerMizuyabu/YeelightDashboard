import {createSelector} from '@ngrx/store';
import {State} from '../reducers/index';
import { LightsState } from '../reducers/lights.reducer';

export const selectLightsState = (state: State) => state.lights

export const selectAllLights = createSelector(
  selectLightsState,
  (lights: LightsState) => Object.keys(lights.entities).map(k => lights.entities[k])
);

export const selectLightById = createSelector(
  selectLights,
  (lights: LightsState, lightId: string) => lights.entities[lightId]
);
