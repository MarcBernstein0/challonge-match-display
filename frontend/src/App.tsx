import { Grid } from '@mui/material';
import { useEffect, useState } from 'react';
import { Match } from './api/api';
import LoadingAnimation from './components/loading';
import CustomizedTables from './components/table';
import { Matches } from './models/matches.interface';
import moment from 'moment';
import { AxiosError } from 'axios';
import { BrowserView } from 'react-device-detect';
import Footer from './components/footer';



function App() {
  const [isLoaded, setIsLoaded] = useState<boolean>(false);
  const [matchResult, setMatches] = useState<Matches[]>([]);

  useEffect(() => {
    const parsedDate = moment();

    console.log(process.env);
    Match.getMatchesV2(parsedDate)
      .then((data) => {
        console.log("data result came back");
        setIsLoaded(true);
        setMatches(data);
      })
      .catch((err: AxiosError) => {
        console.log("error occurred");
        setIsLoaded(true);
        console.error("error occurred on website startup:", err)
      });

    const interval = setInterval(() => {
      Match.getMatchesV2(parsedDate)
        .then((data) => {
          console.log("data result came back");
          setIsLoaded(true);
          setMatches(data);
        })
        .catch((err: AxiosError) => {
          setIsLoaded(true);
          console.error("error occurred on website update:", err)
        });
    }, 30000);
    return () => clearInterval(interval);


  }, []);


  return (
    <span>
      <BrowserView>
        {isLoaded ? (
          <span>
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
                  <Grid item sm={
                    matchResult.length <= 2 ? (12 / matchResult.length) : 4
                  }>
                    <CustomizedTables matchData={game} />
                  </Grid>
                ))}
              <Grid item xs={12}>
                <h1 style={{
                  textAlign: "center"
                }}>
                  IF YOU SEE YOUR MATCH DISPLAYED PLEASE GO AHEAD AND PLAY YOUR MATCH. <br />
                  MAKE SURE TO REPORT RESULTS TO THE TOURNAMENT'S TO.
                </h1>
              </Grid>
            </Grid>
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
      </BrowserView>
      {/* <MobileView>
        <Box display="flex"
             justifyContent="center"
             alignItems="center"
             marginTop="2%"
        >
          <CustomizedMenus games={matchResult} />
        </Box>
      </MobileView>  */}
      <Footer />
    </span>


  );

}

export default App;
