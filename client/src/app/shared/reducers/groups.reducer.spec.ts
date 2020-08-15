import {groupsReducer} from './groups.reducer';
import { Action } from '@ngrx/store';

describe('GroupsReducer', () => {
  it('should return an initialized map', () => {
    const action = {} as Action
    const state = groupsReducer(undefined, action);

    expect(state.ids).toEqual(new Array<string|number>());
    expect(state.entities).toEqual({});
  });
});