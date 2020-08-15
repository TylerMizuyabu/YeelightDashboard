import { createAction, props } from '@ngrx/store'

export const newGroup = createAction(
  '[Groups] New Group',
  props<{groupName: string}>()
)