import { createAction, props } from '@ngrx/store'
import {Yeelight} from '../models/light';

export const addLights = createAction(
  '[Lights Page] addLights',
  props<{lights: Yeelight[]}>()
)