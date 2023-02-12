import { Grid } from '@mui/material';
import moment from 'moment';
import { useEffect, useState } from 'react';
import { Match } from './api/api';
import LoadingAnimation from './components/loading';
import CustomizedTables from './components/table';
import { Matches } from './models/matches.interface';


function App() {
  const [isLoaded, setIsLoaded] = useState<boolean>(false);
  const [matchResult, setMatches] = useState<Matches[]>([]);
  const [isError, setIsError] = useState<boolean>(false);
  const [error, setError] = useState<string>("");

  useEffect(() => {
    console.log(process.env);
    const parsedDate = moment();
    Match.getMatches(parsedDate)
      .then((data) => {
        console.log("data result", data)
        setIsLoaded(true);
        setMatches(data);
      })
      .catch((err) => {
        console.log("error occured");
        setIsLoaded(true);
        setIsError(true);
        setError(err.message);
      });

    const interval = setInterval(() => {
      Match.getMatches(parsedDate)
        .then((data) => {
          console.log(data)
          setIsLoaded(true);
          setMatches(data);
        })
        .catch((err) => {
          setIsLoaded(true);
          setIsError(true);
          setError(err);
        });
    }, 180000);
    return () => clearInterval(interval);
  }, []);


  return (
    <span>
      {isLoaded ? (
        <span>
          {!isError ? (
            <Grid container spacing={3}>
              {matchResult.length === 0 ?
                <Grid item xs={12}>
                  <h1 style={{
                    textAlign: "center"
                  }}>
                    No tournaments
                  </h1>
                </Grid>
                : matchResult.map((game) => (
                  <Grid item xs={
                    matchResult.length <= 2 ? (12/matchResult.length) : 4
                  }>
                    <CustomizedTables matchData={game} />
                  </Grid>
                ))}
                <Grid item xs={12}>
                  <h1 style={{
                    textAlign: "center"
                  }}>
                    IF YOU SEE YOUR MATCH DISPLAYED PLEASE GO AHEAD AND PLAY YOUR MATCH. <br/>
                    MAKE SURE TO REPORT RESULTS TO THE TOURNAMENT'S TO.
                  </h1>
                </Grid>
            </Grid>
          ) : <div>{error}</div>}
        </span>
      )
        :
        <Grid display="flex"
          justifyContent="center"
          alignItems="center"
          minHeight="100vh"
        >
          <LoadingAnimation />
        </Grid>}
    </span>
  );

}

export default App;
