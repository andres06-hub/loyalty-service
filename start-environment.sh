#!/bin/bash

if [ $1 == "upb" ]; then
  echo "Executing docker-compose up --build on the current working directory..."
  docker-compose up --build
  echo "Done"
  exit
fi

if [ $1 == "up" ]; then
  echo "Executing docker-compose up on the current working directory..."
  docker-compose up
  echo "Done"
  exit
fi

if [ $1 == "down" ]; then
  echo "Executing docker-compose down using the current working directory..."
  docker-compose down
  echo "Done"
  exit
fi

echo "Unrecognized command. Please try again using \"up"\ to run the docker compose or \"down"\ if you hace already ran the compose earlier to end the compose process"