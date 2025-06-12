import React from 'react';
import { AppBar, Toolbar, Typography, Container, Button } from '@mui/material';
import Link from 'next/link';

const Layout = ({ children }) => {
  return (
    <>
      <AppBar position="static">
        <Toolbar>
          <Typography variant="h6" sx={{ flexGrow: 1 }}>
            Examen Fullstack - Catálogo & Órdenes
          </Typography>
          <Link href="/" passHref>
            <Button color="inherit">Productos</Button>
          </Link>
          <Link href="/order" passHref>
            <Button color="inherit">Crear Orden</Button>
          </Link>
        </Toolbar>
      </AppBar>
      <Container sx={{ mt: 4 }}>
        {children}
      </Container>
    </>
  );
};

export default Layout;
