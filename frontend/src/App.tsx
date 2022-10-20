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
    const interval = setInterval(() => {
      const parsedDate = moment("2022-08-20", "YYYY-MM-DD");
      Match.getMatches(parsedDate)
        .then((data) => {
          setIsLoaded(true);
          setMatches(data);
        })
        .catch((err) => {
          setIsLoaded(true);
          setIsError(true);
          setError(err);
        });
    }, 5000);
    return () => clearInterval(interval);
  }, []);

  return (
    <span>
      {isLoaded ? (
        <span>
          {!isError ? (
            <Grid container spacing={3}>
              {matchResult.map((game) => (
                <Grid item xs={4}>
                  <CustomizedTables matchData={game} />
                </Grid>
              ))}
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
