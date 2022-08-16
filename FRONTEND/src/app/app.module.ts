import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppComponent } from './app.component';
import { TacheAfficherComponent } from './tache-afficher/tache-afficher.component';

@NgModule({
  declarations: [
    AppComponent,
    TacheAfficherComponent
  ],
  imports: [
    BrowserModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
