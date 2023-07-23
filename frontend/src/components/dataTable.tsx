import { TableComponents } from 'react-virtuoso';
import { Match, Matches } from '../models/matches.interface';
import React from 'react';
import { Paper, Table, TableBody, TableCell, TableContainer, TableHead, TableRow } from '@mui/material';

const VirtuosoTableComponents: TableComponents<Match> = {
  Scroller: React.forwardRef<HTMLDivElement>((props, ref) => (
    <TableContainer component={Paper} {...props} ref={ref} />
  )),
  Table: (props) => (
    <Table {...props} sx={{ borderCollapse: 'separate', tableLayout: 'fixed' }} />
  ),
  TableHead,
  TableRow: ({ item: _item, ...props }) => <TableRow {...props} />,
  TableBody: React.forwardRef<HTMLTableSectionElement>((props, ref) => (
    <TableBody {...props} ref={ref} />
  )),
};

function fixedGameNameContent(gameName: string) {
  return (
    <TableRow>
      <TableCell>
      gameName
      </TableCell>
    </TableRow>
  );
}

function fixedHeaderContent() {
  return (
    <TableRow>
      <TableCell>Match</TableCell>
      <TableCell>Round</TableCell>
      <TableCell>Underway</TableCell>
    </TableRow>
  );
}

export default function ReactVirtualizedTable() {
  return (
    <Paper style={{ height: 400, width: '100%' }}>

    </Paper>
  );
}