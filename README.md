whereisbot
==========

A simple REST server that returns a location for the current system date, intended to be used as a Slack integration.

# Build

Build the image locally.

```
docker build -t ssk2/whereisbot .
```

# Setup

This application assumes a JSON file sitting on an HTTP server.

Here's an example JSON file:
```json
{
  "people": [
    {
      "name": "Peter",
      "locations": [
        {
          "date": "2015-10-13",
          "location": "working from New York"
        },
        {
          "date": "2015-10-14",
          "location": "flying to London"
        }
      ]
    },
    {
      "name": "Paul",
      "locations": [
        {
          "date": "2015-10-13",
          "location": "on vacation in Jakarta"
        },
        {
          "date": "2015-10-14",
          "location": "flying to Tokyo"
        }
      ]
    }
  ]
}
```

# Run

Run the Docker container locally
```
docker run -p 80:8080 -e SOURCE_JSON=https://path/to/dates.json ssk2/whereisbot
```

# Test

Visit [http://localhost/whereis](http://localhost/whereis) to see the output. If no location is defined for that date, it will return `undefined`. If a location is found, it will return the value of the location field.

# Deploy

Push the image up to the internets (assuming a DockerHub account):
```sh
docker push ssk2/whereisbot
```

Use the included Marathon application definition to deploy to Marathon, making sure to update the environment variables as necessary:
```sh
dcos marathon app add marathon.json
```

# Slack Integration

See the instructions [here](https://api.slack.com/slash-commands).

# TODO:

+ Add Google Calendar Integration
+ Add some kind of exception catching
