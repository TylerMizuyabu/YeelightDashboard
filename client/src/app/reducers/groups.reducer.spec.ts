import {groupsReducer} from './groups.reducer';
import { Action } from '@ngrx/store';
import { Yeelight } from '../models/light';

describe('GroupsReducer', () => {
  it('should return an initialized map', () => {
    const action = {} as Action
    const state = groupsReducer(undefined, action)

    expect(state).toEqual(new Map<string, Set<Yeelight>>())
  })
})