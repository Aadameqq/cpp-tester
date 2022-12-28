#include<bits/stdc++.h>

using namespace std;


class UtilMethods{

    public:
     static int randomNum(int mn,int mx){
        return (rand()%mx)+mn;
    }

};

int main(){
    ios_base::sync_with_stdio(0);
    cin.tie(0);
    cout.tie(0);

    int rnd;
    cin>>rnd;
    
    srand(time(NULL)*rnd);

    // Under no circumstances should you edit the lines above //
    // Your code //
    

    int a = UtilMethods::randomNum(1,70);
    int b = UtilMethods::randomNum(a+5,99);

    cout<<a<<" "<<b;



    
}