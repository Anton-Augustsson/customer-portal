import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http'

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent {
  constructor(private http: HttpClient) {}

  ngOnInit(): void {
    console.log('ngOnInit')
    this.http
      .get('https://www.google.com/')
      .subscribe(res => {
        console.log('res', res)
      })
  }
}
