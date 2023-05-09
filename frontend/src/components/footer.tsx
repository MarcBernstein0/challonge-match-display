import { ReactElement } from "react";
import "../css/footer.css"
import { Box, Link, Typography } from "@mui/material";
import { GitHub, Twitter } from "@mui/icons-material";
import { Icon } from '@iconify/react';

export default function Footer(): ReactElement {
  return (
    <Box sx={{
      position: "fixed",
      bottom: 0,
      width: "100%"
    }}>
      <footer className="footer">
        <div className="footer-container">
          <div className="item1">
            <Typography color="white" variant="h5">
              Pendending Matches
            </Typography>
          </div>


          <div className="item2">
            <span style={{ paddingLeft: 5 }}>
              Developed By Marc Bernstein <Link href="https://twitter.com/ujustgotbernied" target="_blank">@KosherSalt</Link>
            </span>
          </div>
          <a
            href="https://github.com/MarcBernstein0/challonge-match-display"
            target="_blank"
            className="item3"
          >
            <GitHub />
          </a>
          <a
            href="https://twitter.com/TravelingCtrlr"
            target="_blank"
            className="item4"
          >
            <Twitter />
          </a>
          <a
            href="https://www.youtube.com/"
            target="_blank"
            className="item5"
          >
            <Icon icon="mdi:twitch" />
          </a>

        </div>
      </footer>
    </Box>
  );
}