import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
const configUrl = 'assets/config.json';

@Injectable()
export class ConfigService {
  constructor(private http: HttpClient) { }
  configUrl = 'assets/config.json';

  getConfig() {
    return this.http.get<any>(this.configUrl);
  }
}