import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { FormComponent } from './pages/form/form.component';
import { IndexComponent } from './pages/index/index.component';
import { IndexDefaultComponent } from './pages/index/index-default/index-default.component';
import { IndexRokComponent } from './pages/index/index-rok/index-rok.component';
import { ServerTypeComponent } from './pages/index/index-default/server-type/server-type.component';

@NgModule({
  declarations: [
    AppComponent,
    FormComponent,
    IndexComponent,
    IndexDefaultComponent,
    IndexRokComponent,
    ServerTypeComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
