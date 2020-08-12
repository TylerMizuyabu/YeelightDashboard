import { createReducer, on } from '@ngrx/store';
import { addLights } from '../actions/add-lights.action';
import { Yeelight } from '../models/light';

export const lightsReducer = createReducer(
  new Set<Yeelight>(),
  on(addLights, (state, {lights}) => (new Set<Yeelight>([...state, ...lights])))
);