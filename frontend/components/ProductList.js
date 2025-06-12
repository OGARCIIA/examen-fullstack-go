import React, { useEffect, useState } from 'react';
import axios from 'axios';
import {
  Card, CardContent, Typography, Grid, CircularProgress
} from '@mui/material';

const ProductList = ({ token }) => {
  const [products, setProducts] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchProducts = async () => {
      try {
        const res = await axios.get('http://localhost:8080/api/products', {
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        setProducts(res.data);
      } catch (err) {
        console.error(err);
      } finally {
        setLoading(false);
      }
    };

    fetchProducts();
  }, [token]);

  if (loading) return <CircularProgress sx={{ display: 'block', mx: 'auto', my: 4 }} />;

  return (
    <Grid container spacing={3} justifyContent="center">
      {products.map((product) => (
        <Grid item xs={12} sm={6} md={4} lg={3} key={product.id}>
          <Card elevation={4} sx={{ minHeight: 180 }}>
            <CardContent>
              <Typography variant="h6" gutterBottom>{product.name}</Typography>
              <Typography variant="body2" color="text.secondary" gutterBottom>
                {product.description}
              </Typography>
              <Typography variant="body1">Precio: ${product.price}</Typography>
              <Typography variant="body1">Stock: {product.stock}</Typography>
            </CardContent>
          </Card>
        </Grid>
      ))}
    </Grid>
  );
};

export default ProductList;
