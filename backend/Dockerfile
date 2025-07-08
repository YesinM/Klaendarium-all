#komentarze są tylko moim konspektem, dlatego można usunąć :)
#from język:vers-"system" obrazu
FROM golang:latest

#folder *wirtualny*
WORKDIR /app
#kopiowanie plików do utworzonego folderu
COPY . . 
#RUN polecenie
RUN go mod download 
RUN go install github.com/air-verse/air@latest
#/backend -copy to> /app
#
EXPOSE 5172

#RUN działa tylko podczas ładowania kontenera, dlatego RUN go run main.go stworzy "martwy" kontener
CMD ["air"]