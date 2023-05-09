import { Box, Grid } from '@mui/material';
import { useEffect, useState } from 'react';
import { Match } from './api/api';
import LoadingAnimation from './components/loading';
import CustomizedTables from './components/table';
import { Matches } from './models/matches.interface';
import moment from 'moment';
import { AxiosError } from 'axios';
import { BrowserView, MobileView } from 'react-device-detect';
import CustomizedMenus from './components/dropdown';
import Footer from './components/footer';



function App() {
  const [isLoaded, setIsLoaded] = useState<boolean>(false);
  const [matchResult, setMatches] = useState<Matches[]>([]);

  console.log(process.env.NODE_ENV);

  useEffect(() => {
    const parsedDate = moment();

    console.log(process.env);
    Match.getMatches(parsedDate)
      .then((data) => {
        console.log("data result came back");
        setIsLoaded(true);
        setMatches(data);
      })
      .catch((err: AxiosError) => {
        console.log("error occured");
        setIsLoaded(true);
        console.error("error occured on website startup:", err)
      });

    const interval = setInterval(() => {
      Match.getMatches(parsedDate)
        .then((data) => {
          console.log("data result came back");
          setIsLoaded(true);
          setMatches(data);
        })
        .catch((err: AxiosError) => {
          setIsLoaded(true);
          console.error("error occured on website update:", err)
        });
    }, 90000);
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
                  <Grid item xs={
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
      <MobileView>
        <Box display="flex"
             justifyContent="center"
             alignItems="center"
             marginTop="2%"
        >
          <CustomizedMenus games={matchResult} />
        </Box>
      </MobileView>
      <Footer />
    </span>


  );

}

export default App;
