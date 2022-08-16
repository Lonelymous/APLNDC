import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TacheAfficherComponent } from './tache-afficher.component';

describe('TacheAfficherComponent', () => {
  let component: TacheAfficherComponent;
  let fixture: ComponentFixture<TacheAfficherComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ TacheAfficherComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(TacheAfficherComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
