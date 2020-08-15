import { Spectator, createComponentFactory } from '@ngneat/spectator';

import { LightListComponent } from './light-list.component';

describe('LightListComponent', () => {
  let spectator: Spectator<LightListComponent>;
  const createComponent = createComponentFactory(LightListComponent);

  it('should create', () => {
    spectator = createComponent();

    expect(spectator.component).toBeTruthy();
  });
});
