import { Injectable } from "@angular/core";
import { MacroEntry } from "../models/macro_entry";
import { User } from "../models/user";

@Injectable()
export class ApiService {
    constructor() {}

    listMacroEntries(userID: number): Promise<MacroEntry[]> {
        return fetch(`http://macronomicon-docker-api-1/macro_entries?userID=${userID}`).then(response => response.json() as Promise<MacroEntry[]>);
    }

    listUsers(): Promise<User[]> {
        return fetch('http://macronomicon-docker-api-1:8080/users').then(response => response.json() as Promise<User[]>);
    }
}