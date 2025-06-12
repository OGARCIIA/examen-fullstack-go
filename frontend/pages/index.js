import React, { useEffect, useState } from 'react';
import axios from 'axios';
import Layout from '../components/Layout';
import ProductList from '../components/ProductList';
import { Typography, Box } from '@mui/material';

const IndexPage = () => {
  const [token, setToken] = useState('');

  useEffect(() => {
    const getToken = async () => {
      try {
        const res = await axios.get('http://localhost:8080/generate-token');
        setToken(res.data.token);
      } catch (err) {
        console.error(err);
      }
    };

    getToken();
  }, []);

  if (!token) return <div>Generando token...</div>;

  return (
    <Layout>
      <Box sx={{ textAlign: 'center', mb: 4 }}>
        <Typography variant="h4" component="h1" gutterBottom>
          Productos
        </Typography>
      </Box>
      <ProductList token={token} />
    </Layout>
  );
};

export default IndexPage;
