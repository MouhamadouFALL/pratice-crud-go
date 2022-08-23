import { Component, ElementRef, EventEmitter, OnInit, Output, ViewChild } from '@angular/core';
import { filter, map, Observable } from "rxjs";
import { ProductService } from "../product.service";
import { Product } from '../product';
import { Router } from '@angular/router';

@Component({
  selector: 'app-list-products',
  templateUrl: './list-products.component.html',
  styleUrls: ['./list-products.component.css']
})
export class ListProductsComponent implements OnInit {

  //@ViewChild('areaSearch', {static: false}) 
  //areaSearch: ElementRef = {} as ElementRef;

  //@Output()
  //sendSearch: EventEmitter<string> = new EventEmitter();

  idcategory: any;


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

  updateProduct(id: string) {
    this.router.navigate(['edit', id]);
  }

  productDetails(id: string){
    this.router.navigate(['details', id]);
  }

  search(){
    //const data = this.areaSearch.nativeElement.value;

    //return this.products = this.products.pipe(filter(item => item.id_category === data));

    //console.log(this.products = this.products.pipe(filter(pro => pro.id_category === data)));
  }

}
