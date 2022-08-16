import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Message } from '@zeropsboost/api-interfaces';

@Component({
  selector: 'zeropsboost-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
})
export class AppComponent {
  hello$ = this.http.get<Message>(
    'https://api-20195-3333.app.zerops.io/api/hello'
  );
  constructor(private http: HttpClient) {}
}
