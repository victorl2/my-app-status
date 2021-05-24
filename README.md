# MyAppStatus
MyAppStatus is a verifier to continually check the **integrity of applications**, the app sends http request in fixed time intervals to update the integrity status for a group of applications; it is possible to defined a group of apps providing the service name and its environments information _( environment name and url )_. This app is created to be included with your custom brand, you can customize a logo and small description that will be displayed in the index page.

## Setup

First you need to configure all settings for your use case, all configurations can be found inside the `config.json` in the root directory of MyAppStatus.

```json
{
    "interval_check": 1, //interval in minutes  between status checks on the services
    "advanced_check": false, //still not used
    "logo_src_url": "https://i.ibb.co/k392Q2M/logo.gif", //
    "logo_url": "http://google.com", //redirect url when the user clicks on the logo
    "subtitle": "Sample subtitle to describe your application", //text used to describe the application in the index page
    "environments": [ // names for the environments used
        "Development",
        "QA",
        "Production"
    ],
    "services": [ // applications that will be verified
        {
            "name": "A Custom API", // name for the application
            "environments": [ //url for health check endpoints ( must respect the order defined before for the environments)
                {
                    "url": "https://google.com"
                },
                {
                    "url": "https://google.com"
                },
                {
                    "url": "https://google.com"
                }
            ]
        },
        ...
    ]
}
```

# Execution

After appying your configuration you can run the application using docker, just run `docker build -t appstatus .` to build the application, the app will build and create a executable inside itself, its not required to have golang installed in your local machine; after the build you can run `docker run -ti appstatus -p 8080:8080`, now just acess the your localhost:8080 to see the app in action.
