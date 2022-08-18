import { Component, OnInit } from '@angular/core';
import { map, Observable } from "rxjs";
import { ProductService } from "../product.service";
import { Product } from '../product';
import { Router } from '@angular/router';

@Component({
  selector: 'app-list-products',
  templateUrl: './list-products.component.html',
  styleUrls: ['./list-products.component.css']
})
export class ListProductsComponent implements OnInit {

  product: Product = new Product();
  products = new Observable<any>();

  constructor(private productService: ProductService, 
    private router: Router) { }

  ngOnInit(): void {
    this.reloadData();
  }

  reloadData(){
    this.products = this.productService.getProductList();
  }

  deleteProduct(id: number) {
    this.productService.deleteProduct(id);
    this.reloadData();
  }

  updateProduct(id: number) {
    this.router.navigate(['edit', id]);
  }

  productDetails(id: number){
    this.router.navigate(['details', id]);
  }

}
