# myapp-go
ESD Projekt

# start Project 
cd myapp-go/src/teamstar

$ docker compose build

$ docker compose up

# Use the REST API
login admin:

curl -H "Content-Type: application/json" -d '{"id":1,"name":"admin","role":"trainer"}' localhost:8080/login


create user (as trainer):

curl -H "Content-Type: application/json" -d '{"name":"marc","role":"player"}' localhost:8080/trainer/users


Liste aller User:

curl localhost:8080/userlist


login player (oben erstellt):

curl -H "Content-Type: application/json" -d '{"id":1,"name":"admin","role":"trainer"}' localhost:8080/login


create training (als trainer):

curl -H "Content-Type: application/json" -d '{"topic":"Regeneration","content":"entspannter Lauf und Fußballtennis. Laufschuhe mitbringen!","date":"10.07.2022","user":{"name":"admin","role":"trainer","id":1}}' localhost:8080/trainer/training 


create user (als trainer):

curl -H "Content-Type: application/json" -d '{"name":"marc","role":"player"}' localhost:8080/trainer/users



Liste aller Trainings:

curl localhost:8080/player/trainings


get Training mit ID=1:

curl localhost:8080/player/trainings/1


create Feedback Ja (als spieler):

curl -H "Content-Type: application/json" -d '{"status":"YES","reason":"","user":{"name":"marc","role":"player","id":2}}' localhost:8080/player/trainings/1/feedback


create Feedback Nein mit Begründung (als spieler):

curl -H "Content-Type: application/json" -d '{"status":"NO","reason":"Muss lernen","user":{"name":"marc","role":"player","id":2}}' localhost:8080/player/trainings/1/feedback


Feedback von Training 1 anschauen:

curl localhost:8080/player/trainings/1


DELETE Training mit ID=1:

curl -X DELETE localhost:8080/trainer/trainings/1