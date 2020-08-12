import { createAction, props } from '@ngrx/store'

export const newGroup = createAction(
  '[Groups Page] newGroup',
  props<{groupName: string}>()
)