import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { Tâche, PrioritéEnum } from './tâche';

@Injectable({
  providedIn: 'root'
})
export class TasksService {

  constructor() { }

  obtenirDesTâches(): Observable<Tâche[]> {
    let lisTâches: Tâche[] = [];

    var myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/json");
    myHeaders.append("Access-Control-Allow-Origin", "*");

    var requestOptions = {
      method: 'GET',
      headers: myHeaders,
    };

    fetch("http://81.182.202.18:4000/tasks", requestOptions)
      .then(response => {
        response.json().then(x => {
          let values: any[] = x;

          for (let i = 0; i < values.length; i++) {
            let obj = values[i];
            let tâcheNouveau: Tâche = {
              TâcheIdentificationNuméro: obj.TareaIdentificaciónNúmero,
              UtilisateurIdentificationNuméro: obj.UsuarioIdentificaciónNúmero,
              Nom: obj.Nombre,
              Description: obj.Descripción,
              Priorité: obj.Prioridad,
              Fini: obj.Hecho,
              Terme: obj.Plazo,
              CrééEn: obj.CreadoEn
            }
            lisTâches.push(tâcheNouveau);
          }
        })
      })
      .then(result => console.log(result))
      .catch(error => console.log('error', error));

    return of(lisTâches);
  }
}
