# perfin-api

This project was generated with tooling created by Ambassador Labs

## Inspiration

The PerFin (Personal Finance) Calculator API was inspired by my constant need for a doing financial calculations. "How much will I get if this is compounded?", "What's my financial independence number?" and other questions that come up as I have gotten more interested in personal finance. I need a simple, reliable, and readily accessible way to perform common financial calculations. This API is the answer and can be integrated into any other application or used from the command line like I plan to do.

## What it does

The PerFin Calculator API provides a set of endpoints for performing various financial calculations, including simple and compound interest, loan payments, future and present value, inflation-adjusted returns, fixed deposit maturity calculations and investment ROI. It accepts requests and returns the calculated results in JSON format.

## How we built it

I built the PerFin Calculator API using Ambassador Blackbird, it started out with a swagger yaml file that was created with the swagger editor. This file was fed to the blackbird cli with which an api was created, mocked and code was generated. I then edited and debugged the generated golang code to complete the api and deployed it.

## Challenges we ran into

Some of the challenges I faced were:
- Currently the code generation done by Blackbird is in Golang, I only have passing knowledge in it so it was an interesting learning experience.
- I had originally wanted to do the api in python with FastAPI but I could not figure out how to create and deploy an api with python. (Though I think I have figured it out)

## Accomplishments that we're proud of

I'm proud of creating an API from swagger yaml file to being completely deployed using blackbird. I have an API created, working and deployed essentially within two days. 
Blackbird made this very easy to deploy, I did not have to mess around with any servers, configurations or any of that. And anyone can test the application as well, super proud of that.

## What we learned

I learned the following:

- Blackbird is really easy to use, to generate code from a yaml file
- I learned how to create a golang api and how that's structured

## What's next for PerFin Calculator API

Future enhancements could include:

* **More complex calculations:**  Adding support for more advanced financial calculations, I have a few more calculations I use every once in a while but also these could be used more by other people.
* **User authentication and authorization:** Allowing users to create accounts and store their financial data securely. 
* **Securing the API** Adding an api key with blackbird to secure it.
* **Some additional features**  Maybe pulling country specific financial information, like inflation, interest rates and more.

