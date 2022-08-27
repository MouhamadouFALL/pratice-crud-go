import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Product } from './product';

@Injectable({
  providedIn: 'root'
})
export class ProductService {
  headers = { 'content-type': 'application/json' };

  private baseUrl: string = '/api/products';

  constructor(private http: HttpClient) { }

  getProduct(id: number): Observable<Product> {
    return this.http.get<Product>(`${this.baseUrl}/${id}`);
  }

  createProduct(product: Product): Observable<Product> {
    return this.http.post<Product>(`${this.baseUrl}`, product);
  }

  updateProduct(id: number, value: Product): Observable<Product> {
    return this.http.put<Product>(`${this.baseUrl}/${id}`, value);
  }

  deleteProduct(id: number): Observable<any>{
    return this.http.delete<any>(`${this.baseUrl}/${id}`, { headers: this.headers });
  }

  getProductList(): Observable<Product[]> {
    return this.http.get<Product[]>(`${this.baseUrl}`);
  }

}
