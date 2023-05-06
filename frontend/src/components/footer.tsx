import { GitHub, Twitter } from "@mui/icons-material";
import { Box, Container, Grid, Link, Typography } from "@mui/material";
import { ReactElement } from "react";

// const footerItems = [
//   {
//     link: "https://github.com/MarcBernstein0/challonge-match-display",
//     icon: <GitHub />
//   },
//   {
//     link: "https://twitter.com/TravelingCtrlr",
//     icon: <Twitter />
//   },
// ]

export default function Footer(): ReactElement {
  return (
    <Box
      sx={{
        width: "100%",
        height: "auto",
        backgroundColor: "secondary.main",
        paddingTop: "1rem",
        paddingBottom: "1rem",
      }}
    >
      <Container maxWidth="lg">
        <Grid container
          direction="row"
          justifyContent="center"
          alignItems="center">
          <Grid item xs={12} sx={{
              textAlign: "center",
              border: 1
            }}>
            <Typography color="black" variant="h5">
              Pendending Matches
            </Typography>
          </Grid>
          <Grid item xs={6} sx={{ 
            border: 1,
            textAlign: "center" 
            }}>
            <p>Developed by Marc Bernstein(KosherSalt)</p><Link href="https://github.com/MarcBernstein0/challonge-match-display">
              <GitHub />
            </Link>
          </Grid>
          {/* {footerItems.map(item => (
            <Grid item xs={6} sx={{
              border: 1
            }}>
              <Link href={item.link} >
                {item.icon}
              </Link>
            </Grid>
          ))} */}

        </Grid>
      </Container>
    </Box>
  )
}