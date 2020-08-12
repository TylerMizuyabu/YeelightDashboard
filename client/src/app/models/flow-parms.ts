
export interface FlowParams {
  count: number;
  action: FlowAction;
  tuples: FlowTuple[];
}

export enum FlowAction {
  RecoverToPreviousState = 0,
  KeepCurrentState,
  TurnOff
}

export interface FlowTuple {
  duration: number;
  mode: number;
  value: number;
  brightness: number;
}