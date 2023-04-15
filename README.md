# Predict CO2 Emissions

What if we could combine the power of machine learning and the speed of Go?

In this practice i will to predict CO2 Emissions using a multiple linear regression model and implement the model in Go.

$$y = \theta_0 + \theta_1 x_1 + \theta_2 x_2 + \cdots + \theta_n x_n$$

First i will analyze the data, get some insights and get the best features to use in the model. I will build the model in a jupyter notebook to analyze the results and compare function loss and accuracy. 

Finally i will implement the model in Go thru a REST API using Gin.

This experiment will be helpful for a current project that i'm working on.

## TODO

- Add an explanation of each column of the dataframe in `analyzing-data.ipynb`
- Update README.md in main directory:
    - Add some chart-images from any notebooks of the project.
    - Add some code blocks in markdown of how to make the http request and how does the endpoint will look.
    - Add instructions to reproduce Go-API using the docker-compose file.