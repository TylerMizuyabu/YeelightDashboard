import {lightsReducer} from './lights.reducer';
import { Action } from '@ngrx/store';

describe('LightsReducer', () => {
  it('should return an initialized map', () => {
    const action = {} as Action
    const state = lightsReducer(undefined, action);

    expect(state.ids).toEqual(new Array<string|number>());
    expect(state.entities).toEqual({});
  });

});