import { createReducer, on } from '@ngrx/store';
import { addLights } from '../actions/add-lights.action';
import { createEntityAdapter, EntityState } from '@ngrx/entity';
import { Yeelight } from '../models/light';

export interface LightsState extends EntityState<Yeelight> {
}

const adapterLights = createEntityAdapter<Yeelight>({
  selectId: (light: Yeelight) => (light.id),
});

export const initialLightsState = adapterLights.getInitialState({});

export const lightsReducer = createReducer(
  initialLightsState,
  on(addLights, (state, { lights }) => (adapterLights.upsertMany(lights, state)))
);