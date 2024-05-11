import { HttpClientModule } from '@angular/common/http';
import { importProvidersFrom } from '@angular/core';

export const appConfig = {
  providers: [
    importProvidersFrom(HttpClientModule),
  ]
};