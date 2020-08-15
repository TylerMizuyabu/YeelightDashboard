import { Yeelight } from './light';

export interface Group {
  name: string;
  lights: Set<Yeelight>;
}