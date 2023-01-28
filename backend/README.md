# match-display
Display upcoming matches from multiple brackets running on challonge.com

# TODO:
1) New backend flow to add caching and lessen api calls
    1) Call backend to get tournaments and there participants, add a timestamp to the results.
        2) Every time the api is called check the timestamp and see if it has exceeded a set time. If it has, call challonge to see if any new tournaments are running or if any have ended. If new tournament exists then add to existing tournaments list and get participants, if not then do nothing. 
    2) Use the tournament and participant data to get pending matches. Setup frontend to call this route every x min