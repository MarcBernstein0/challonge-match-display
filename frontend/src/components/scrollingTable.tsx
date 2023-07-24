import * as React from 'react';
import Paper from '@mui/material/Paper';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TablePagination from '@mui/material/TablePagination';
import TableRow from '@mui/material/TableRow';
import { Match, Matches } from '../models/matches.interface';

interface Column {
  id: 'match' | 'round' | 'underway';
  label: string;
  minWidth?: number;
  align?: 'right';
  format?: (value: number) => string;
}

const columns: readonly Column[] = [
  { id: 'match', label: 'Name', minWidth: 170 },
  { id: 'round', label: 'ISO\u00a0Code', minWidth: 100 },
  {
    id: 'underway',
    label: 'Population',
    minWidth: 170,
    align: 'right',
    format: (value: number) => value.toLocaleString('en-US'),
  },
];

interface TableProps {
  matchData: Matches;
}

interface Data {
  match: string;
  round: string;
  underway: boolean;
}

function createData(
  matchData: Match
): Data {
  const match = `${matchData.player1_name} vs ${matchData.player2_name}`;
  const round = matchData.round <= -1 ? `winners ${Math.abs(matchData.round)}` : `losers ${Math.abs(matchData.round)}`;
  const underway = matchData.underway;
  return { match, round, underway };
}

// const rows = [
//   createData('India', 'IN', 1324171354),
//   createData('China', 'CN', 1403500365),
//   createData('Italy', 'IT', 60483973),
//   createData('United States', 'US', 327167434),
//   createData('Canada', 'CA', 37602103),
//   createData('Australia', 'AU', 25475400),
//   createData('Germany', 'DE', 83019200),
//   createData('Ireland', 'IE', 4857000),
//   createData('Mexico', 'MX', 126577691),
//   createData('Japan', 'JP', 126317000),
//   createData('France', 'FR', 67022000),
//   createData('United Kingdom', 'GB', 67545757),
//   createData('Russia', 'RU', 146793744),
//   createData('Nigeria', 'NG', 200962417),
//   createData('Brazil', 'BR', 210147125),
// ];

function createRows(matchList: Match[]): Data[] {
  const res: Data[] = [];
  for(const match of matchList){
    res.push(createData(match));
  }
  return res;
}

export default function StickyHeadTable({ matchData }: TableProps) {
  const [page, setPage] = React.useState(0);
  const [rowsPerPage, setRowsPerPage] = React.useState(10);

  const handleChangePage = (event: unknown, newPage: number) => {
    setPage(newPage);
  };

  const handleChangeRowsPerPage = (event: React.ChangeEvent<HTMLInputElement>) => {
    setRowsPerPage(+event.target.value);
    setPage(0);
  };

  const matches: Match[] = matchData.match_list;
  const rows: Data[] = createRows(matches);

  // return (
  //   <Paper sx={{ width: '100%', overflow: 'hidden' }}>
  //     <TableContainer sx={{ maxHeight: 440 }}>
  //       <Table stickyHeader aria-label="sticky table">
  //         <TableHead>
  //           <TableRow>
  //             {columns.map((column) => (
  //               <TableCell
  //                 key={column.id}
  //                 align={column.align}
  //                 style={{ minWidth: column.minWidth }}
  //               >
  //                 {column.label}
  //               </TableCell>
  //             ))}
  //           </TableRow>
  //         </TableHead>
  //         <TableBody>
  //           {matches
  //             .slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage)
  //             .map((match) => {
  //               return (
  //                 <TableRow hover role="checkbox" tabIndex={-1} key={match.player1_id}>
  //                   {columns.map((column) => {
  //                     const value = match.;
  //                     return (
  //                       <TableCell key={column.id} align={column.align}>
  //                         {column.format && typeof value === 'number'
  //                           ? column.format(value)
  //                           : value}
  //                       </TableCell>
  //                     );
  //                   })}
  //                 </TableRow>
  //               );
  //             })}
  //         </TableBody>
  //       </Table>
  //     </TableContainer>
  //     <TablePagination
  //       rowsPerPageOptions={[10, 25, 100]}
  //       component="div"
  //       count={rows.length}
  //       rowsPerPage={rowsPerPage}
  //       page={page}
  //       onPageChange={handleChangePage}
  //       onRowsPerPageChange={handleChangeRowsPerPage}
  //     />
  //   </Paper>
  // );
}