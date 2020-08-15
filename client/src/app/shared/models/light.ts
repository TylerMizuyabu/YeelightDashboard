import { FlowParams } from "./flow-parms";

export interface Yeelight {
  id: string;
  model: LightModel;
  isOn: boolean;
  brightness: number;
  mode: LightMode;
  ct: number;
  rgb: number;
  hue: number;
  sat: number;
  name: string;
  flowing: FlowMode;
  flowParameters: FlowParams;
}

export type LightModel = 'mono' | 'color' | 'stripe' | 'ceiling' | 'bslamp';

export enum LightMode {
  DefaultLightMode = 0,
  ColorMode,
  ColorTemperatureMode,
  HSVMode,
  ColorFlowMode,
  NightLightMode
}

export enum FlowMode {
  NoFlow = 0,
  FlowRunning
}