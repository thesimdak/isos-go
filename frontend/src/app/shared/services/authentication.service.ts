import { Injectable } from "@angular/core";
@Injectable({
  providedIn: "root",
})
export class AuthenticationService {
  public heslo: String = "palestra";

  public authenticated: Boolean = false;

  constructor() {}

  public authenticate(psw: String): void {
    if (this.heslo === psw) {
      this.authenticated = true;
    }
  }

  public isAuthenticated() {
    return this.authenticated;
  }
}
