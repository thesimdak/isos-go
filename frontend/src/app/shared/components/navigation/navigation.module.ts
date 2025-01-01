import { NgModule } from '@angular/core';
import { NavigationComponent } from './navigation.component';
import { CommonModule } from '@angular/common';
import { BrowserModule } from '@angular/platform-browser';
import { AppRoutingModule } from '../../../app-routing.module';

@NgModule({
    imports: [
      AppRoutingModule,
      BrowserModule,
      CommonModule
              ],
  declarations: [
    NavigationComponent
  ],
  providers: [],
  exports: [NavigationComponent],
})
export class NavigationModule { }