## Instructions for running on dev-environment

- Run VSCode command `Remote-Containers: Reopen in Container`
- Container configuration is in `.devcontainer` folder. 
    - Name of container is `data-analysis` by default.
- In case of installing a new package
    - run `pip install <package-name>` in the terminal
    - then run `pip freeze > requirements.txt` to update the requirements file.

## TODO

- Add an explanation of each column of the dataframe in `analyzing-data.ipynb`